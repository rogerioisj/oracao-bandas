const openRegisterModalButton = document.querySelector('#open-band-modal');
const openMobileRegisterModalButton = document.querySelector('#mobile-open-band-modal');
const registerModal = document.querySelector('#band-modal');
const closeRegisterModalButton = document.querySelector('#close-band-modal');

openRegisterModalButton.addEventListener('click', () => {
    registerModal.showModal();
});

closeRegisterModalButton.addEventListener('click', () => {
    registerModal.close();
});

openMobileRegisterModalButton.addEventListener('click', () => {
    registerModal.showModal();
});

const openUserModalButton = document.querySelector('#open-user-modal');
const mobileOpenUserModalButton = document.querySelector('#mobile-open-user-modal');
const UserModal = document.querySelector('#user-modal');
const closeUserModalButton = document.querySelector('#close-user-modal');

openUserModalButton.addEventListener('click', () => {
    UserModal.showModal();
});

mobileOpenUserModalButton.addEventListener('click', () => {
    UserModal.showModal();
});

closeUserModalButton.addEventListener('click', () => {
    UserModal.close();
});

const openUpdatedUserModalButton = document.querySelector('#open-update-user-modal');
const mobileOpenUpdatedUserModalButton = document.querySelector('#mobile-open-update-user-modal');
const UpdateUserModal = document.querySelector('#update-user-modal');
const closeUpdateUserModalButton = document.querySelector('#close-update-user-modal');

openUpdatedUserModalButton.addEventListener('click', () => {
    UpdateUserModal.showModal();
});

mobileOpenUpdatedUserModalButton.addEventListener('click', () => {
    UpdateUserModal.showModal();
});

closeUpdateUserModalButton.addEventListener('click', () => {
    UpdateUserModal.close();
});