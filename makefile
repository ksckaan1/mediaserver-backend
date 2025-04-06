.PHONY: gen-proto
gen-proto:
	protoc \
		--go_out=Mprotofiles/media_service.proto=shared/pb/mediapb:. \
		--go-grpc_out=Mprotofiles/media_service.proto=shared/pb/mediapb:. \
		protofiles/media_service.proto
	protoc \
		--go_out=Mprotofiles/movie_service.proto=shared/pb/moviepb:. \
		--go-grpc_out=Mprotofiles/movie_service.proto=shared/pb/moviepb:. \
		--proto_path=protofiles \
		protofiles/movie_service.proto
	protoc \
		--go_out=Mprotofiles/tmdb_service.proto=shared/pb/tmdbpb:. \
		--go-grpc_out=Mprotofiles/tmdb_service.proto=shared/pb/tmdbpb:. \
		--proto_path=protofiles \
		protofiles/tmdb_service.proto
	protoc \
		--go_out=Mprotofiles/series_service.proto=shared/pb/seriespb:. \
		--go-grpc_out=Mprotofiles/series_service.proto=shared/pb/seriespb:. \
		--proto_path=protofiles \
		protofiles/series_service.proto
	protoc \
		--go_out=Mprotofiles/season_service.proto=shared/pb/seasonpb:. \
		--go-grpc_out=Mprotofiles/season_service.proto=shared/pb/seasonpb:. \
		--proto_path=protofiles \
		protofiles/season_service.proto
	protoc \
		--go_out=Mprotofiles/episode_service.proto=shared/pb/episodepb:. \
		--go-grpc_out=Mprotofiles/episode_service.proto=shared/pb/episodepb:. \
		--proto_path=protofiles \
		protofiles/episode_service.proto
	protoc \
		--go_out=Mprotofiles/user_service.proto=shared/pb/userpb:. \
		--go-grpc_out=Mprotofiles/user_service.proto=shared/pb/userpb:. \
		--proto_path=protofiles \
		protofiles/user_service.proto
	protoc \
		--go_out=Mprotofiles/auth_service.proto=shared/pb/authpb:. \
		--go-grpc_out=Mprotofiles/auth_service.proto=shared/pb/authpb:. \
		--proto_path=protofiles \
		protofiles/auth_service.proto
	