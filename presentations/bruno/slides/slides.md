---
theme: ../../../theme
title: We don't talk about Bruno
highlighter: shiki
drawings:
  enabled: false
lineNumbers: true
colorSchema: light
hideInToc: true
mdc: true

themeConfig:
  eventLogo: 'https://bdxio.fr/favicon.ico'
  eventUrl: 'https://bdxio.fr/'
  bluesky: 'imflog'
  blueskyUrl: 'https://bsky.app/profile/imflog.bsky.social'

transition: slide-left
layout: intro
---

# Let's talk about Bruno

<div class="w-xs ma relative" style="z-index: 1">
    <img class="rounded-lg shadow-lg object-cover z-10" src="/images/bruno_labrador_transparent.png" alt="Bruno" />
</div>

---
layout: presenter
presenterImage: /images/florian.png
---

# Florian Garcia

Tech lead at <a  href="https://www.synapse-medicine.com/">Synapse Medicine</a>

- Bordeaux 🍷, France 🥖
- Focusing on improving healthcare with technology ⚕️
- I often write stuff in my personal <a href="https://imflog.github.io/">blog</a>
- Say hi at <a href="https://bsky.app/profile/imflog.bsky.social"><logos-bluesky mr-1 />@imflog.bsky.social</a>

[We are looking for new talents](https://careers.synapse-medicine.com/fr/jobs) ! 🚀

<!--
We are looking for SRE and Product Owner notably.
-->

---
layout: text-image
media: /images/UI-screenshot.png
---

# What is Bruno ?

* An open source API Client 🤖
* Like Postman <logos-postman-icon mr-1 />, Insomnia <logos-insomnia mr-1 /> ...
* Fast ⚡ and Git friendly <logos-git-icon mr-1 />
* Offline only, no cloud sync ! 
* Written in Javascript <logos-javascript mr-1 />
* UI, CLI and VS Code extension <logos-visual-studio-code mr-1 />

<div class="ma" style="width:10rem">
    <img src="/images/logo-bruno.png"/>
</div>

<!--
API client qui permet aussi de tester.
Git friendly -> les fichiers sont stockés en plain-text sur le disque.
Offline only -> pas de synchronisation cloud 
-->

---
layout: new-section
---

# Demo

<div class="w-xs ma">
    <img src="/images/bruno-no-no-no.gif" width="400" />
</div>

---
layout: default
zoom: 0.8
---

## VS Postman

<br/>

|                        | **Postman**                                                                                   | **Bruno**                              |
|------------------------|-----------------------------------------------------------------------------------------------|----------------------------------------|
| **Request protocols**  | HTTP, GraphQL, gRPC, WebSocket, Socket.IO, MQTT, SOAP.                                        | HTTP and GraphQL.                      | 
| **Collection format**  | Open source JSON format.                                                                      | Plain text markup language called Bru. |
| **Collection storage** | Proprietary cloud                                                                             | In your file system.                   |
| **Collection run**     | Free and basic plan at 25 runs per month. Professional at 250 runs, unlimited for Enterprise. | Unlimited runs, anywhere.              |
| **Collaboration**      | Paid feature to share with more than 3 users.                                                 | Free, just use git                     |

---
layout: default
zoom: 0.8
---

## VS Postman

<br />

|                                 | **Postman**                                                                                                                                                              | **Bruno**                                                      |
|---------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|
| **API requests transfert**      | Through Postman servers in the web app. <br/> [Some reports](https://community.postman.com/t/working-in-offline-mode/20174/51) of proxied requests from the desktop app. | Directly from your computer.                                   |
| **Offline ?**                   | [No](https://community.postman.com/t/working-in-offline-mode/20174/37)                                                                                                   | Yes.                                                           |
| **Scripting with NPM modules**  | Using workarounds like pulling from a CDN or copying the library code in a variable.                                                                                     | Use a package.json, like usual                                 |
| **CLI**                         | Postman CLI or Newman.<br/> Both have limits on Postman server API calls amount based on purchased plan.                                                                 | [Bruno CLI](https://docs.usebruno.com/bru-cli/overview). Free. |
| **Local performance testing ?** | Yes with run and user limits.                                                                                                                                            | Not yet.<br/> Planned in the golden edition.                   |
| **Additional features**         | [Many](https://www.postman.com/features/)                                                                                                                                | [Not so many](https://docs.usebruno.com/).                     |

---
layout: two-cols
---

## Pricing ?

![PostmanPricing](/images/pricing-postman.png)

::right::

![BrunoPricing](/images/pricing-bruno.png)

<!--
Il y a aussi une version pour une utilisation personnelle qui coûte 19E en one-time payment pour 2 ordi et 2 ans de maj.
-->

---
class: 'grid text-center align-self-center justify-self-center'
layout: default
---

# Thank you 

<div align="center">
    <a href="https://github.com/ImFlog/presentations">https://github.com/imflog/presentations</a>
    <img src="/images/qrcode-github.png" width="300"/>
</div>
