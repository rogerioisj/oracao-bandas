/*NAV BAR */

const currentPath = window.location.pathname;
const homeLink = document.getElementById('home-link');
const userLink = document.getElementById('users-link');

console.log("PATH: ", currentPath)
console.log("HOME: ", homeLink)
console.log("USERS: ", userLink)

if (currentPath === '/' && homeLink) {
    homeLink.classList.add('border-r-4');
    homeLink.classList.add('border-red-600');
    homeLink.classList.remove('border-white');
}

if (currentPath === '/' && userLink && homeLink) {
    homeLink.classList.add('border-r-4');
    homeLink.classList.add('border-red-600');
    homeLink.classList.remove('border-white');

    userLink.classList.remove('border-r-4');
    userLink.classList.remove('border-red-600');
    userLink.classList.add('border-white');
}

if (currentPath !== '/' && userLink && homeLink) {
    userLink.classList.add('border-r-4');
    userLink.classList.add('border-red-600');
    userLink.classList.remove('border-white');

    homeLink.classList.remove('border-r-4');
    homeLink.classList.remove('border-red-600');
    homeLink.classList.add('border-white');
}

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

if (openLoginModalButton !== null) {
    openLoginModalButton.addEventListener('click', () => {
        loginModal.showModal();
    });

    closeLoginModalButton.addEventListener('click', () => {
        loginModal.close();
    });
}
