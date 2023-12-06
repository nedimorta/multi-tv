import React, { useState, useEffect } from 'react';
import ChannelForm from './components/ChannelForm';
import ChannelList from './components/ChannelList';
import ChannelAdjuster from './components/ChannelAdjuster';

interface Channel {
  name: string;
  id: string;
}

const App: React.FC = () => {
  const [channels, setChannels] = useState<Channel[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/channels')
      .then(response => response.json())
      .then(data => setChannels(data))
      .catch(error => console.error('Error fetching channels:', error));
  }, []);

  const addChannel = (channel: Channel) => {
    fetch('http://localhost:8080/api/channels', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(channel),
    })
    .then(response => response.json())
    .then(newChannel => setChannels([...channels, newChannel]))
    .catch(error => console.error('Error adding channel:', error));
  };

  const onAdjustChannelCount = (newCount: number) => {
    // Implement the logic to update the channel count on the backend
    // This might involve an API call to the backend
  };

  return (
    <div>
      <h1>YouTube Channel Manager</h1>
      <ChannelForm onAddChannel={addChannel} />
      <ChannelList channels={channels} />
      <ChannelAdjuster currentCount={channels.length} onAdjust={onAdjustChannelCount} />
    </div>
  );
};

export default App;
