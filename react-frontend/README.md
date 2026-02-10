# Overview

This is a client web application that interacts with a golang [backend](vscode-webview://1mtbqp480jbktthmnhtasivcq8o4dp3ns99qlg0mt23ne4gsj5rg/index.html?id=a8a74031-7793-4b09-9447-28adbc9e404b&parentId=3&origin=78708ab4-7bee-458a-b02f-a1498c30e496&swVersion=4&extensionId=patmood.rich-markdown-editor&platform=electron&vscode-resource-base-authority=vscode-resource.vscode-cdn.net&parentOrigin=vscode-file://vscode-app) for open authorization and session management. The app itself comes with two pages that allow the user to login and logout.


> This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).


## Features

* React Context - streamlined state management for tracking user authorization
* App.js - defines a client-based routing system using React Router
* Login.js - user interface for initiating an OAuth workflow with Google handle by sending request to the golang backend
* Dashboard.js - displays the user’s Identity Provider data such as email, name, and avatar photo


## Available Scripts

In the project directory, you can run: `npm start`


This runs the app in the development mode.Open <http://localhost:8080> to view it in your browser. The page will reload when you make changes.You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.It correctly bundles React in production mode and optimizes the build for the best performance.



