// Node: 6
// Path: /console/src/components/BidPanel.tsx
// Purpose: UI input + bid history panel

import React from 'react';

const BidPanel: React.FC = () => {
  return (
    <div>
      <input type="text" placeholder="Enter bid" />
      <button>Submit Bid</button>
      <div>
        <h3>Bid History</h3>
        <ul>
          <li>$100.00</li>
          <li>$90.00</li>
          <li>$80.00</li>
        </ul>
      </div>
    </div>
  );
};

export default BidPanel;