/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../src/static/**/*.{html,js}", "../src/views/partials/*.html", "../src/views/home/*.html", "../src/views/auth/*.html"],
  theme: {
    extend: {},
    fontFamily: {
      body: ["'Protest Guerrilla', sans-serif"]
    },
  },
  plugins: [],
}

