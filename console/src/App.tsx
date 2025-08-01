// Node: 4 (related)
// Path: /console/src/App.tsx
// Purpose: Root component for the frontend application.

import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import AuctionPage from './pages/auction/[id]';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/auction/:id" element={<AuctionPage />} />
      </Routes>
    </Router>
  );
}

export default App;
