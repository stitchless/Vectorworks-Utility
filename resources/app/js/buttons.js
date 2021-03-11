let buttons = {
    init: () => {
        // Edit Serial Button Elements
        const buttonElements = document.querySelectorAll('table td i.fal')

        // Add button attributes to button elements
        buttonElements.forEach(i => i.addEventListener('click', buttons.editSerial))
        buttonElements.forEach(i => i.addEventListener('mouseover', btnMouseOver))
        buttonElements.forEach(i => i.addEventListener('mouseout', btnMouseOut))
    },
    editSerial: (element) => {
        const software = element.target.dataset.software
        const year = element.target.dataset.year
        index.editSerial(software, year)
    }
}

function btnMouseOver() {
    this.style.cursor = "pointer"
    this.classList.remove('fal')
    this.classList.add('far')
}

function btnMouseOut() {
    this.style.cursor = "auto"
    this.classList.remove('far')
    this.classList.add('fal')
}