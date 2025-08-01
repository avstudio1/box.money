// Node: 4
// Path: /console/src/pages/auction/[id].tsx
// Purpose: Dynamic auction page: Zustand, UI, stream

import React from 'react';
import { useParams } from 'react-router-dom';
import { useAuctionStream } from '../../hooks/useAuctionStream';
import BidPanel from '../../components/BidPanel';

const AuctionPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const { bidHistory, connectionStatus, submitBid } = useAuctionStream(id || '');

  return (
    <div>
      <h1>Auction: {id}</h1>
      <p>Connection Status: {connectionStatus}</p>
      <BidPanel />
      {/* Display bid history */}
    </div>
  );
};

export default AuctionPage;