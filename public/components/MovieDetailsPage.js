import API from '../services/API.js'

export class MovieDetailsPage extends HTMLElement {
  id = null
  movie = null

  async render() {
    try {
      this.movie = await API.getMovieById(this.movie)
    } catch (error) {
      // TODO: Alert the user
      console.log(error)
      return
    }
    const template = document.getElementById('template-movie-details')
    const content = template.content.cloneNode(true)
    this.appendChild(content)

    this.querySelector('h2').textContent = this.movie.title
    this.querySelector('h3').textContent = this.movie.tagline
  }

  // By default connectedCallback is not an async function
  connectedCallback() {
    this.movie = 14 // TODO - get the movie id
    this.render()
  }
}

customElements.define('movie-details-page', MovieDetailsPage)
