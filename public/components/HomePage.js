export class HomePage extends HTMLElement {
  connectedCallback() {
    // template is like a class, it is not the element
    const template = document.getElementById('template-home')
    // creates the element - like an instance of a class
    const content = template.content.cloneNode(true)
    this.appendChild(content)
  }
}

// slashes are mandatory for customElements
// 'homepage' won't work
customElements.define('home-page', HomePage)
