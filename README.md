# Portfolio/Blog

A custom blog application written with Go, SQL, and CSS. It doubles as my portfolio and I self-host it on a low-power mini PC in my basement. I use Caddy as a reverse proxy and a bash script to manage deployments. I have a webhook setup to run my deployment script whenever there is a push to main on this repo.

## Background
I’ve been using Obsidian to organize my notes. Since Obsidian just uses Markdown files on my local machine, I decided to start publishing some of my notes as a blog. I set up a VPS and used Deno to create a blog with minimal setup, deploying it from my Obsidian vault using a webhook. However, after encountering issues following a Deno update, I decided to rebuild the blog in Go. I absolutely love Go and now I try to write as much Go as I can. If you're interested in more details, you can [read more here](https://dominicgerman.com/posts/4).
