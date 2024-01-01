/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./src/**/*.templ"],
  darkMode: 'class',
  theme: {
    fontFamily: {
      crenzo: 'Crenzo',
      baunk: 'Baunk',
      triakis: 'Triakis'
    }
  },
  plugins: [
    require('./rx-colors')
  ],
}
