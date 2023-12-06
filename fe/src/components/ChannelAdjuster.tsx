import React, { useState } from 'react';

const ChannelAdjuster: React.FC<{ currentCount: number, onAdjust: (newCount: number) => void }> = ({ currentCount, onAdjust }) => {
    const [channelCount, setChannelCount] = useState<number>(currentCount);

    const handleAdjustment = () => {
        onAdjust(channelCount);
    };

    return (
        <div>
            <input type="number" value={channelCount} onChange={(e) => setChannelCount(parseInt(e.target.value))} />
            <button onClick={handleAdjustment}>Adjust Channels</button>
        </div>
    );
};

export default ChannelAdjuster;
