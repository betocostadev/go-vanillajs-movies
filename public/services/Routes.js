import { HomePage } from '../components/HomePage'
import { MovieDetailsPage } from '../components/MovieDetailsPage'
import { MoviePage } from '../components/MoviePage'

// Use regexes to generate dynamic routes
export const routes = [
  {
    path: '/',
    component: HomePage,
  },
  {
    path: '/movies/',
    component: MoviePage,
  },
  {
    path: /\/movies\/(\d+)/, // /movies/14
    component: MovieDetailsPage,
  },
  {
    path: '/account/register',
    component: RegisterPage,
  },
  {
    path: '/account/login',
    component: LoginPage,
  },
  {
    path: '/account/',
    component: AccountPage,
  },
  {
    path: '/account/favorites',
    component: FavoritesPage,
  },
  {
    path: '/account/watchlist',
    component: WatchlistPage,
  },
]
