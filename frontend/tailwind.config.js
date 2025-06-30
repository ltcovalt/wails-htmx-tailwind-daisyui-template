/** @type 
 {import('tailwindcss').Config}
 */
export default {
	content: [
		"./index.html",
		"../templates/**/*.{html}",
		"./src/**/*.{js,ts,jsx,tsx}",
	],
	theme: {
		extend: {},
	},
	plugins: [],
}