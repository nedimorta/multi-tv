import React, { useState } from 'react';

interface ChannelFormProps {
  onAddChannel: (channel: { name: string; id: string }) => void;
}

const ChannelForm: React.FC<ChannelFormProps> = ({ onAddChannel }) => {
  const [name, setName] = useState<string>('');
  const [id, setId] = useState<string>('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onAddChannel({ name, id });
    setName('');
    setId('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" value={name} onChange={e => setName(e.target.value)} placeholder="Channel Name" required />
      <input type="text" value={id} onChange={e => setId(e.target.value)} placeholder="Channel ID" required />
      <button type="submit">Add Channel</button>
    </form>
  );
};

export default ChannelForm;
