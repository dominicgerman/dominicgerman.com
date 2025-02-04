# Portfolio/Blog

A custom blog application written with Go, SQL, and CSS. It doubles as my portfolio and I self-host it on a Raspberry Pi in my basement. I use Caddy as a reverse proxy and Ansible to manage deployments. 

## Background
I’ve been using Obsidian to organize my notes. Since Obsidian just uses Markdown files on my local machine, I decided to start publishing some of my notes as a blog. I set up a VPS and used Deno to create a blog with minimal setup, deploying it from my Obsidian vault using a webhook. However, after encountering issues following a Deno update, I decided to rebuild the blog in Go. I absolutely love Go and now I try to write as much Go as I can. 

## Usage
Although possible, this app isn't really designed to be used by anyone but me. The UI would have to be completely re-written. But it does have forms for user authentication, creating posts, updating posts, and an endpoint for deleting posts. It has middleware for handling CSRF protection, session management and logging. If you're interested in the details, you can [read more here](https://dominicgerman.com/posts/4).
