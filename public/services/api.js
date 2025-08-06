export const API = {
  baseURL: '/api/',
  getTopMovies: async () => {
    // const response = await fetch('/api/movies/top')
    // const result = await response.json()
    return await API.fetch('movies/top')
  },
  getRandomMovies: async () => {
    return await API.fetch('movies/random')
  },
  getMovieById: async (id) => {
    return await API.fetch(`movies/${id}`)
  },
  searchMovies: async (query, order, genre) => {
    return await API.fetch(`/movies/search`, { query, order, genre })
  },
  getGenres: async () => {
    return await API.fetch('genres')
  },
  fetch: async (service, args) => {
    try {
      const queryString = args ? new URLSearchParams(args).toString() : ''
      const response = await fetch(API.baseURL + service + '?' + queryString)
      const result = await response.json()
      return result
    } catch (e) {
      console.error(e)
      app.showError()
    }
  },
}

export default API
