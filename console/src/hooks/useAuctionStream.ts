// Node: 5
// Path: /console/src/hooks/useAuctionStream.ts
// Purpose: Bid/Result streaming over gRPC-Web

import { useState, useEffect } from 'react';
import { grpc } from '@improbable-eng/grpc-web';
import { AuctionServiceClient } from '../gen/auction/Auction_serviceServiceClientPb';
import { BidRequest, BidUpdate } from '../gen/auction/auction_pb';

export const useAuctionStream = (auctionId: string) => {
  const [bidHistory, setBidHistory] = useState<BidUpdate[]>([]);
  const [connectionStatus, setConnectionStatus] = useState('disconnected');

  useEffect(() => {
    const client = new AuctionServiceClient('http://localhost:5173');
    const stream = client.bidStream();

    stream.on('data', (response) => {
      setBidHistory((prev) => [...prev, response]);
    });

    stream.on('status', (status) => {
      setConnectionStatus(grpc.Code[status.code]);
    });

    stream.on('end', () => {
      setConnectionStatus('disconnected');
    });

    return () => {
      stream.cancel();
    };
  }, [auctionId]);

  const submitBid = (amount: number) => {
    const client = new AuctionServiceClient('http://localhost:5173');
    const request = new BidRequest();
    request.setAuctionId(auctionId);
    request.setAmount(amount);
    // How to get user ID?
    // request.setUserId('user-id'); 

    const stream = client.bidStream();
    stream.write(request);
  };

  return { bidHistory, connectionStatus, submitBid };
};