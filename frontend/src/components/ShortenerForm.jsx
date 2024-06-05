import React, { useState } from 'react';

function ShortenerForm() {
  const [originalURL, setOriginalURL] = useState('');
  const [shortenedURL, setShortenedURL] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    const response = await fetch('/api/shorten', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ original_url: originalURL })
    });
    const result = await response.json();
    setShortenedURL(result.short_url);
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="url"
          value={originalURL}
          onChange={(e) => setOriginalURL(e.target.value)}
          placeholder="Enter your URL here"
          required
        />
        <button type="submit">Shorten</button>
      </form>
      {shortenedURL && (
        <div>
          Shortened URL: <a href={shortenedURL}>{shortenedURL}</a>
        </div>
      )}
    </div>
  );
}

export default ShortenerForm;
