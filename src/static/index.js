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

const openRegisterModalButton = document.querySelector('#open-band-modal');
const registerModal = document.querySelector('#band-modal');
const closeRegisterModalButton = document.querySelector('#close-band-modal');

openRegisterModalButton.addEventListener('click', () => {
    registerModal.showModal();
});

closeRegisterModalButton.addEventListener('click', () => {
    registerModal.close();
});

const openUserModalButton = document.querySelector('#open-user-modal');
const UserModal = document.querySelector('#user-modal');
const closeUserModalButton = document.querySelector('#close-user-modal');

openUserModalButton.addEventListener('click', () => {
    UserModal.showModal();
});

closeUserModalButton.addEventListener('click', () => {
    UserModal.close();
});
