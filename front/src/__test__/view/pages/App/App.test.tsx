import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import App from '../../../../view/pages/App/App';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/your IPv4/i);
  expect(linkElement).toBeInTheDocument();
});
