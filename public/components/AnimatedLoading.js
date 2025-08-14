class AnimatedLoading extends HTMLElement {
  constructor() {
    super()
  }
  connectedCallback() {
    // grabs the dataset of elements. Like data-element, data-testid
    const elements = this.dataset.elements
    const width = this.dataset.width
    const height = this.dataset.height

    for (let i = 0; i < elements; i++) {
      const wrapper = document.createElement('div')
      wrapper.classList.add('loading-wave')
      wrapper.style.width = width
      wrapper.style.height = height
      wrapper.style.margin = '10px'
      wrapper.style.display = 'inline-block'
      this.appendChild(wrapper)
    }
  }
}

customElements.define('animated-loading', AnimatedLoading)
