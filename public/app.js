import { API } from './services/API.js'
// import { HomePage } from './components/HomePage.js'

import './components/YoutubeEmbed.js'
import './components/AnimatedLoading.js'
import { Router } from './services/Router.js'
// console.log(API)
window.addEventListener('DOMContentLoaded', () => {
  app.Router.init()
  // document.querySelector('main').appendChild(new HomePage())
})

window.app = {
  api: API,
  Router,
  showError: (message = 'There was an error.', goToHome = true) => {
    document.getElementById('alert-modal').showModal()
    document.querySelector('#alert-modal p').textContent = message
    if (goToHome) app.Router.go('/')
  },
  closeError: () => {
    document.getElementById('alert-modal').close()
  },
  search: (event) => {
    event.preventDefault()
    const query = document.querySelector('input[type=search]').value
    if (query.length > 1) {
      app.Router.go(`/movies/?query=${query}`)
    }
  },
}
