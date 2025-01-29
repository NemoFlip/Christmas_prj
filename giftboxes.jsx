import React, { useEffect, useState } from 'react';
import './styles/boxes.css';

const ProductGrid = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/users/data');
        if (!response.ok) {
          throw new Error('Ошибка загрузки данных');
        }
        const data = await response.json();

        const adaptedProducts = data.map((item, index) => ({
          id: index,
          name: item.gift_name,
          description: item.gift_description,
          image: item.image_url || 'https://via.placeholder.com/150',
          price: item.price ? `${item.price} руб.` : 'Цена не указана',
        }));

        setProducts(adaptedProducts);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchProducts();
  }, []);

  if (loading) {
    return <p>Загрузка...</p>;
  }

  if (error) {
    return <p className="error">Ошибка: {error}</p>;
  }

  return (
    <div className="product-section">
      <h2 id="selection-title" className="selection-title">Selection based on your data</h2>
      <div className="grid">
        {products.map((product) => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
};

const ProductCard = ({ product }) => {
  return (
    <div className="card">
      <img src={product.image} alt={product.name} className="card-image" />
      <h3>{product.name}</h3>
      <p>{product.description}</p>
      <p>{product.price}</p>
    </div>
  );
};

export default ProductGrid;
