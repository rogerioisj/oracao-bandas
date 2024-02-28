const burger = document.querySelector('#burger');
const menu = document.querySelector('#menu');

const showMenu = () => {
    if (menu.classList.contains('hidden')) {
        menu.classList.remove('hidden');
    } else {
        menu.classList.add('hidden');
    }
};

burger.addEventListener('click', showMenu);
