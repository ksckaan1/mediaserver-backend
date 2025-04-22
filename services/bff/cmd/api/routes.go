package main

import (
	"bff-service/internal/core/app/auth"
	"bff-service/internal/core/app/episode"
	"bff-service/internal/core/app/media"
	"bff-service/internal/core/app/movie"
	"bff-service/internal/core/app/season"
	"bff-service/internal/core/app/series"
	"bff-service/internal/core/app/user"
	"shared/enums/usertype"
	"shared/password"
	"shared/pb/authpb"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"
	"shared/pb/userpb"

	"github.com/gofiber/fiber/v2"
)

type UserTypeMWFunc = func(userType usertype.UserType) func(c *fiber.Ctx) error

func initAuthRoutes(authClient authpb.AuthServiceClient, userClient userpb.UserServiceClient) (*fiber.App, fiber.Handler, UserTypeMWFunc) {
	app := fiber.New()
	authMW := auth.NewAuthMiddleware(authClient, userClient).Handle
	userTypeMW := auth.NewUserTypeMiddleware().RequiredUserType

	app.Post("/login", auth.NewLogin(authClient).Handle)
	app.Get("/logout", authMW, auth.NewLogout(authClient).Logout)
	return app, authMW, userTypeMW
}

func initUserRoutes(
	userClient userpb.UserServiceClient,
	authMW fiber.Handler,
	userTypeMW UserTypeMWFunc,
) *fiber.App {
	pw := password.New()
	app := fiber.New()
	app.Post("/register", h(user.NewRegister(userClient)))
	app.Use(authMW)
	app.Get("/profile", h(user.NewProfile(userClient)))
	app.Put("/:id/password", h(user.NewUpdatePassword(userClient, pw)))
	app.Put("/:id/user-type", userTypeMW(usertype.Admin), h(user.NewUpdateUserType(userClient)))
	return app
}

func initMediaRoutes(mediaClient mediapb.MediaServiceClient, userTypeMW UserTypeMWFunc) *fiber.App {
	app := fiber.New()
	app.Post("/", userTypeMW(usertype.Admin), h(media.NewCreateMedia(mediaClient)))
	app.Get("/:media_id", h(media.NewGetMediaByID(mediaClient)))
	app.Get("/", h(media.NewListMedias(mediaClient)))
	app.Put("/:media_id", userTypeMW(usertype.Admin), h(media.NewUpdateMediaByID(mediaClient)))
	app.Delete("/:media_id", userTypeMW(usertype.Admin), h(media.NewDeleteMediaByID(mediaClient)))
	return app
}

func initMovieRoutes(movieClient moviepb.MovieServiceClient, userTypeMW UserTypeMWFunc) *fiber.App {
	app := fiber.New()
	app.Post("/", userTypeMW(usertype.Admin), h(movie.NewCreateMovie(movieClient)))
	app.Get("/search", h(movie.NewSearchMovie(movieClient)))
	app.Get("/:movie_id", h(movie.NewGetMovieByID(movieClient)))
	app.Get("/", h(movie.NewListMovies(movieClient)))
	app.Put("/:movie_id", userTypeMW(usertype.Admin), h(movie.NewUpdateMovieByID(movieClient)))
	app.Delete("/:movie_id", userTypeMW(usertype.Admin), h(movie.NewDeleteMovieByID(movieClient)))
	return app
}

func initSeriesRoutes(seriesClient seriespb.SeriesServiceClient, userTypeMW UserTypeMWFunc) *fiber.App {
	app := fiber.New()
	app.Post("/", userTypeMW(usertype.Admin), h(series.NewCreateSeries(seriesClient)))
	app.Get("/search", h(series.NewSearchSeries(seriesClient)))
	app.Get("/:series_id", h(series.NewGetSeriesByID(seriesClient)))
	app.Get("/", h(series.NewListSeries(seriesClient)))
	app.Put("/:series_id", userTypeMW(usertype.Admin), h(series.NewUpdateSeriesByID(seriesClient)))
	app.Delete("/:series_id", userTypeMW(usertype.Admin), h(series.NewDeleteSeriesByID(seriesClient)))
	return app
}

func initSeasonRoutes(seasonClient seasonpb.SeasonServiceClient, userTypeMW UserTypeMWFunc) *fiber.App {
	app := fiber.New()
	app.Post("/", userTypeMW(usertype.Admin), h(season.NewCreateSeason(seasonClient)))
	app.Get("/:season_id", h(season.NewGetSeasonByID(seasonClient)))
	app.Get("/", h(season.NewListSeasons(seasonClient)))
	app.Put("/:season_id", userTypeMW(usertype.Admin), h(season.NewUpdateSeasonByID(seasonClient)))
	app.Delete("/:season_id", userTypeMW(usertype.Admin), h(season.NewDeleteSeasonByID(seasonClient)))
	return app
}

func initEpisodeRoutes(episodeClient episodepb.EpisodeServiceClient, userTypeMW UserTypeMWFunc) *fiber.App {
	app := fiber.New()
	app.Post("/", userTypeMW(usertype.Admin), h(episode.NewCreateEpisode(episodeClient)))
	app.Get("/:episode_id", h(episode.NewGetEpisodeByID(episodeClient)))
	app.Get("/", h(episode.NewListEpisodes(episodeClient)))
	app.Put("/:episode_id", userTypeMW(usertype.Admin), h(episode.NewUpdateEpisodeByID(episodeClient)))
	app.Delete("/:episode_id", userTypeMW(usertype.Admin), h(episode.NewDeleteEpisodeByID(episodeClient)))
	return app
}
