export class MovieItem extends HTMLElement {
  constructor(movie) {
    super()
    this.movie = movie
  }

  connectedCallback() {
    // create the template - in this case we will just create it here instead of importing it
    this.innerHTML = `
      <a href="#">
        <article>
          <img src="${this.movie.poster_url}" alt="${this.movie.title} poster" />
          <p>${this.movie.title} (${this.movie.release_year})</p>
        </article>
      </a>
    `
  }
}

customElements.define('movie-item', MovieItem)
