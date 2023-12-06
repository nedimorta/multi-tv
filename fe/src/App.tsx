import React, { useState, useEffect } from 'react';
import ChannelForm from './components/ChannelForm';
import ChannelList from './components/ChannelList';

interface Channel {
  name: string;
  id: string;
}

const App: React.FC = () => {
  const [channels, setChannels] = useState<Channel[]>([]);

  useEffect(() => {
    // Fetch channels from the backend
    // ...
  }, []);

  const addChannel = (channel: Channel) => {
    // Add channel to the backend
    // ...
  };

  return (
    <div>
      <h1>YouTube Channel Manager</h1>
      <ChannelForm onAddChannel={addChannel} />
      <ChannelList channels={channels} />
    </div>
  );
};

export default App;
