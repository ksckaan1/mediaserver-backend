.PHONY: gen-proto
gen-proto:
	protoc \
		--go_out=Mprotofiles/media_service.proto=common/pb/mediapb:. \
		--go-grpc_out=Mprotofiles/media_service.proto=common/pb/mediapb:. \
		protofiles/media_service.proto
	protoc \
		--go_out=Mprotofiles/movie_service.proto=common/pb/moviepb:. \
		--go-grpc_out=Mprotofiles/movie_service.proto=common/pb/moviepb:. \
		--proto_path=protofiles \
		protofiles/movie_service.proto
	protoc \
		--go_out=Mprotofiles/tmdb_service.proto=common/pb/tmdbpb:. \
		--go-grpc_out=Mprotofiles/tmdb_service.proto=common/pb/tmdbpb:. \
		--proto_path=protofiles \
		protofiles/tmdb_service.proto
	protoc \
		--go_out=Mprotofiles/series_service.proto=common/pb/seriespb:. \
		--go-grpc_out=Mprotofiles/series_service.proto=common/pb/seriespb:. \
		--proto_path=protofiles \
		protofiles/series_service.proto
	