const currentPath = window.location.pathname;
const homeLink = document.getElementById('home-link');
const mobileHomeLink = document.getElementById('mobile-home-link');
const userLink = document.getElementById('users-link');
const mobileUserLink = document.getElementById('mobile-users-link');

console.log("PATH: ", currentPath)
console.log("HOME: ", homeLink)
console.log("USERS: ", userLink)

if (currentPath === '/' && homeLink) {
    homeLink.classList.add('border-r-4');
    homeLink.classList.add('border-red-600');
    homeLink.classList.remove('border-white');
}

if (currentPath === '/' && mobileHomeLink) {
    mobileHomeLink.classList.add('border-r-4');
    mobileHomeLink.classList.add('border-red-600');
    mobileHomeLink.classList.remove('border-white');
}

if (currentPath === '/' && userLink && homeLink) {
    homeLink.classList.add('border-r-4');
    homeLink.classList.add('border-red-600');
    homeLink.classList.remove('border-white');

    userLink.classList.remove('border-r-4');
    userLink.classList.remove('border-red-600');
    userLink.classList.add('border-white');
}

if (currentPath === '/' && mobileUserLink && mobileHomeLink) {
    mobileHomeLink.classList.add('border-r-4');
    mobileHomeLink.classList.add('border-red-600');
    mobileHomeLink.classList.remove('border-white');

    mobileUserLink.classList.remove('border-r-4');
    mobileUserLink.classList.remove('border-red-600');
    mobileUserLink.classList.add('border-white');
}

if (currentPath !== '/' && userLink && homeLink) {
    userLink.classList.add('border-r-4');
    userLink.classList.add('border-red-600');
    userLink.classList.remove('border-white');

    homeLink.classList.remove('border-r-4');
    homeLink.classList.remove('border-red-600');
    homeLink.classList.add('border-white');
}

if (currentPath !== '/' && mobileUserLink && mobileHomeLink) {
    mobileUserLink.classList.add('border-r-4');
    mobileUserLink.classList.add('border-red-600');
    mobileUserLink.classList.remove('border-white');

    mobileHomeLink.classList.remove('border-r-4');
    mobileHomeLink.classList.remove('border-red-600');
    mobileHomeLink.classList.add('border-white');
}

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

const openMobileLoginModalButton = document.querySelector('#mobile-open-login-modal');
const mobileLoginModal = document.querySelector('#login-modal');
const closeMobileLoginModalButton = document.querySelector('#close-login-modal');

if (openMobileLoginModalButton !== null) {
    openMobileLoginModalButton.addEventListener('click', () => {
        mobileLoginModal.showModal();
    });

    closeMobileLoginModalButton.addEventListener('click', () => {
        mobileLoginModal.close();
    });
}

document.addEventListener('DOMContentLoaded', () => {
    const burger = document.getElementById('burger');
    const mobileMenu = document.getElementById('mobile-menu');

    const toggleMobileMenu = () => {
        mobileMenu.classList.toggle('translate-x-full');
        mobileMenu.classList.toggle('translate-x-0');
    };

    burger.addEventListener('click', (event) => {
        event.stopPropagation();
        toggleMobileMenu();
    });

    document.addEventListener('click', (event) => {
        const isClickInsideMenu = mobileMenu.contains(event.target);
        const isClickOnBurger = burger.contains(event.target);

        if (!mobileMenu.classList.contains('translate-x-full') && !isClickInsideMenu && !isClickOnBurger) {
            mobileMenu.classList.add('translate-x-full');
            mobileMenu.classList.remove('translate-x-0');
        }
    });
});
