import { API } from './services/API.js'
import { HomePage } from './components/HomePage.js'
import { MovieDetailsPage } from './components/MovieDetailsPage.js'
import './components/AnimatedLoading.js'
// console.log(API)
window.addEventListener('DOMContentLoaded', () => {
  document.querySelector('main').appendChild(new HomePage())
  document.querySelector('main').appendChild(new MovieDetailsPage())
})

window.app = {
  search: (event) => {
    console.log('Search called')
    event.preventDefault()
    const keywords = document.querySelector('input[type=search]').value
    // Call api, etc
  },
  api: API,
}
