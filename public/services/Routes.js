import { HomePage } from '../components/HomePage.js'
import { LoginPage } from '../components/LoginPage.js'
import { MovieDetailsPage } from '../components/MovieDetailsPage.js'
import { MoviesPage } from '../components/MoviesPage.js'
import { RegisterPage } from '../components/RegisterPage.js'

// Use regexes to generate dynamic routes
export const routes = [
  {
    path: '/',
    component: HomePage,
  },
  {
    path: '/movies/',
    component: MoviesPage,
  },
  {
    // /movies/14
    path: /\/movies\/(\d+)/,
    component: MovieDetailsPage,
  },
  {
    path: '/account/register/',
    component: RegisterPage,
  },
  {
    path: '/account/login/',
    component: LoginPage,
  },
  // {
  //   path: '/account/',
  //   component: AccountPage,
  // },
  // {
  //   path: '/account/favorites',
  //   component: FavoritesPage,
  // },
  // {
  //   path: '/account/watchlist',
  //   component: WatchlistPage,
  // },
]
