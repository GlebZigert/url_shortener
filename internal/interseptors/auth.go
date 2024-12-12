package interseptors

import (
	"context"
	"errors"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor проверяет наличие и валидность токена в заголовках запроса
func (s *Interseptors) AuthInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Получаем заголовки из контекста
	log.Println("--AuthInterceptor")

	//достаем токен из входных данных
	token, err := GetTokenFromCtx(ctx)

	if err != nil {

		return nil, status.Error(codes.InvalidArgument, "")

	}

	log.Println("token: ", token)

	uid, err := s.GetUserID(token)

	if err != nil || uid == 0 {
		log.Println("не смог взять uid из входных")
		uid, err = s.users.CreateNextUID()
		if err != nil {
			s.logger.Error("GetUserID: ", map[string]interface{}{
				"err": err.Error(),
			})

			return nil, status.Error(codes.Internal, "")

		}
		log.Println("next UID: ", uid)

		jwt, err := s.BuildJWTString(uid)
		if err != nil {

			s.logger.Error("BuildJWTString: ", map[string]interface{}{
				"err": err.Error(),
			})
			return nil, status.Error(codes.Internal, "")

		}

		// create and set header

		header := metadata.Pairs("authorisation", jwt)
		grpc.SetHeader(ctx, header)

		ctx = s.SetNewFlag(ctx, true)

	}

	//если нет токена - гененируем его и возвращаем его в ответе  с ошибкой сodes.Unauthenticated
	ctx = s.SetUID(ctx, uid)
	return handler(ctx, req)
}

var NoTokenErr error = errors.New("no valid token")

func GetTokenFromCtx(ctx context.Context) (string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {

		return "", NoTokenErr
		//	return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	authHeader := md.Get("authorisation")
	if len(authHeader) == 0 {

		return "", NoTokenErr
	}
	if len(authHeader[0]) == 0 {

		return "", NoTokenErr
	}

	return authHeader[0], NoTokenErr
}

func SetTokentoCtx(uid int, ctx context.Context) context.Context {

	md := metadata.Pairs("authorisation", strconv.Itoa(uid))
	return metadata.NewOutgoingContext(ctx, md)
}
