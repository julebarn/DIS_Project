import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		host: process.env.HOST || 'localhost',
		port: parseInt(process.env.PORT),
		proxy: {
			'/api': process.env.API_PROXY || 'http://localhost:8080'
		},
	}
});
