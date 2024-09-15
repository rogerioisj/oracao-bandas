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

const openLoginModalButton = document.querySelector('#open-login-modal');
const loginModal = document.querySelector('#login-modal');
const closeLoginModalButton = document.querySelector('#close-login-modal');

openLoginModalButton.addEventListener('click', () => {
    loginModal.showModal();
});

closeLoginModalButton.addEventListener('click', () => {
    loginModal.close();
});
