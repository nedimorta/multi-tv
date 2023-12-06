import React from 'react';

interface Channel {
  name: string;
  id: string;
}

interface ChannelListProps {
  channels: Channel[];
}

const ChannelList: React.FC<ChannelListProps> = ({ channels }) => {
  return (
    <ul>
      {channels.map(channel => (
        <li key={channel.id}>{channel.name}</li>
      ))}
    </ul>
  );
};

export default ChannelList;
