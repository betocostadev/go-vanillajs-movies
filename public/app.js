import { API } from './services'
console.log(API)
window.app = {
  search: (event) => {
    console.log('Search called')
    event.preventDefault()
    const keywords = document.querySelector('input[type=search]').value
    // Call api, etc
  },
}

window.addEventListener('DOMContentLoaded', () => {})
