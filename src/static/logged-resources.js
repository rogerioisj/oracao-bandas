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