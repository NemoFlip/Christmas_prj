import React from 'react';
import './styles/boxes.css'; 

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

const ProductGrid = ({ products }) => {
  if (!Array.isArray(products) || products.length === 0) {
    return <div>Нет данных о продуктах</div>;
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

export default ProductGrid;