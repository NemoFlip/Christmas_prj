import React from 'react';
import { Header } from './components/header'; 
import { FindBlock } from './components/findblock';
import ProductGrid from './components/giftboxes';
import './App.css';

function App() {
  return (
    <>
      <Header /> ,
      <FindBlock />,
      <ProductGrid />
    </>
  );
}

export default App;
