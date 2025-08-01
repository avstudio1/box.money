// Node: 4 (related)
// Path: /console/src/main.tsx
// Purpose: Entrypoint for the React application.

import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)