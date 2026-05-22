import { useState } from "react";
import type { CartItem, Product } from "../types";
import { CartContext } from "./CartContext";

export const CartProvider = ({ children }: { children: React.ReactNode }) => {
  const [cart, setCart] = useState<CartItem[]>([]);

  const addToCart = (product: Product) => {
    setCart((prev) => {
    const existing = prev.find(
      (item) => item.productId === product.id
    );

    if (existing) {
      return prev.map((item) =>
        item.productId === product.id
          ? { ...item, amount: item.amount + 1 }
          : item
      );
    }

    return [...prev, { productId: product.id, amount: 1 }];
  });
  };


  const clearCart = () => setCart([]);

  const getTotal = () => cart.reduce((sum, item) => sum + item.amount, 0);

  return (
    <CartContext.Provider value={{ cart, addToCart, clearCart, getTotal }}>
      {children}
    </CartContext.Provider>
  );
};

