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
        setProducts(data);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    if (products.length === 0) {
      return;
    }

    fetchProducts();
  }, [products]);

  if (loading || error) {
    return null; 
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
      <img
        src={product.image}
        alt={product.name}
        className="card-image"
      />
      <h3>{product.name}</h3>
      <p>Цена: {product.price} руб.</p>
    </div>
  );
};

export default ProductGrid;
