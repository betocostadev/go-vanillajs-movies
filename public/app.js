import { HomePage } from './components/HomePage.js'
import { API } from './services/API.js'
// console.log(API)
window.addEventListener('DOMContentLoaded', () => {
  document.querySelector('main').appendChild(new HomePage())
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
