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
  search: (event) => {
    console.log('Search called')
    event.preventDefault()
    const keywords = document.querySelector('input[type=search]').value
    // Call api, etc
  },
}
