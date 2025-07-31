// Path: /console/src/main.tsx
// Purpose: The main entrypoint for the React frontend application.

import React from 'react';
import ReactDOM from 'react-dom/client';

const App = () => (
  <div style={{ fontFamily: 'sans-serif', padding: '2rem', lineHeight: '1.6' }}>
    <h1>Goldstream Console</h1>
    <p>The frontend service is running correctly.</p>
    <p>The full infrastructure stack is now operational.</p>
  </div>
);


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);