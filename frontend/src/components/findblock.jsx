import React, { useState } from 'react';
import './styles/findblock.css';
import ProductGrid from './giftboxes.jsx';

export function FindBlock() {
  const [recipientText, setRecipientText] = useState('');
  const [showProducts, setShowProducts] = useState(false);
  const [products, setProducts] = useState([]); 

  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/users/data', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ text: recipientText }), 
      });

      if (!response.ok) {
        throw new Error('Ошибка при отправке данных');
      }

      const result = await response.json();
      console.log('Ответ от сервера:', result);

      const adaptedProducts = result.map((item, index) => ({
        id: index,
        name: item.gift_name,
        description: item.gift_description,
        image: item.image_url || 'https://via.placeholder.com/150',
        price: item.price ? `${item.price} руб.` : 'Цена не указана',
      }));

      setProducts(adaptedProducts); 
      setShowProducts(true);
    } catch (error) {
      console.error('Ошибка:', error);
      alert('Ошибка отправки данных. Попробуйте позже.');
    }
  };

  return (
    <div className="main">
      <div className="block">
        <h1 className="findtext">Find the best gift</h1>
        <div className="WAtext">Write about gift recipient</div>
        <textarea
          name="gift recipient"
          className="textfield"
          placeholder="Мама любит готовить, вязать и слушать музыку"
          value={recipientText}
          onChange={(e) => setRecipientText(e.target.value)}
        ></textarea>
        <div>
          <button className="search" onClick={handleSubmit}>
            Search
          </button>
        </div>
      </div>

      {showProducts && <ProductGrid products={products} />} {/* Передаем данные в ProductGrid */}
    </div>
  );
}
