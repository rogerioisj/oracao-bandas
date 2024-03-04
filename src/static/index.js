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

const openModalButton = document.querySelector('#open-login-modal');
const modal = document.querySelector('#login-modal');
const closeModalButton = document.querySelector('#close-login-modal');

openModalButton.addEventListener('click', () => {
    modal.showModal();
});

closeModalButton.addEventListener('click', () => {
    modal.close();
});
