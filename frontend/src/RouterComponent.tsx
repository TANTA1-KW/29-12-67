import React, { StrictMode } from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import SignInPages from "./authentication/Login";
import SignUpPages from "./authentication/Register";

import NavbarFoodService from "./food_service/page/navbar/NavbarFoodService";

import Menu from "./food_service/page/menu/Menu";
import MenuDetail from "./food_service/page/menu_detail/MenuDetail";
import CompletePage from "./stripe/CompletePage";

import { OrderProvider } from "./food_service/context/OrderContext";
import StripeCheckout from "./stripe/StripeCheckout";
import { loadStripe } from "@stripe/stripe-js";
import { Elements } from "@stripe/react-stripe-js";

import OrderSummary from "./food_service/page/order_item/OrderItem";
import Promotion from "./promotion/pages/promotion";
import PromotionCreate from "./promotion/pages/promotion/create";
import PromotionEdit from "./promotion/pages/promotion/edit";
import Review from "./review/pages/review";
// import App from "./stripe2/App";
// import Return from "./stripe2/return";

import Employees from "./employee/pages/";
import EmployeeCreate from "./employee/pages/create";
 import EmployeeEdit from "./employee/pages/edit";

 import Customers from "./customer/pages";
 import CustomerCreate from "./customer/pages/create";
 import CustomerEdit from "./customer/pages/edit/inedex";

const stripePromise = loadStripe("pk_test_51QOxoF4QmAAjQ0QzsimUKy0RcgMxNPvfbmCm6OJurQzEGULD1u2OfTSGfdd0OwpEW0tzpdkQvmQSZKvbq9waUceD00PaT9sjdJ");

const RouterComponent: React.FC = () => {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <SignInPages />,
    },
    {
      path: "/signup",
      element: <SignUpPages />,
    },

    {
      path: "/employee",
      children: [
        {
          path: "/employee",
          element: <Employees />,
        },
        {
          path: "/employee/create",
          element: <EmployeeCreate />,
        },
        {
          path: "/employee/edit/:id",
          element: <EmployeeEdit />,
        },
      ],
    },
    {
      path: "/customer",
      children: [
        {
          path: "/customer",
          element: <Customers />,
        },
        {
          path: "/customer/create",
          element: <CustomerCreate />,
        },
        {
          path: "/customer/edit/:id",
          element: <CustomerEdit />,
        },
     ],
    },
    {
      path: "/review",
      children: [
        {
          path: "/review",
          element: <Review />,
        },
        // {
        //   path: "/review/create",
        //   element: <ReviewCreate />,
        // },
        // {
        //   path: "/review/edit/:id",
        //   element: <ReviewEdit />,
        // },
      ]
    },
    {
      path: "/promotion",
      // element: <Navbar />,
      children: [
        {
          path: "/promotion",
          element: <Promotion />,
        },
        {
          path: "/promotion/create",
          element: <PromotionCreate />,
        },
        {
          path: "/promotion/edit/:id",
          element: <PromotionEdit />,
        },
      ],
    },
    {
      path: "/login/food-service",
      element: <OrderProvider><NavbarFoodService /></OrderProvider>,
      children: [
        {
          path: "order",
          element: <Menu />,
        },
        {
          path: "menu-detail",
          element: <MenuDetail />,
        },
        {
          path: "order-summary",
          element: <OrderSummary />,
        },
        {
          path: "order-summary/checkout",
          element: (    
              <StripeCheckout />
          ),
        },
      ],
    },
    {
      path: "/login/food-service/order/order-summary/checkout/complete",
      element: (
        <OrderProvider>
        <Elements stripe={stripePromise}>
              <CompletePage />
          </Elements>
        </OrderProvider>
      ),
    },
    // {
    //   path:"/checkout", 
    //   element: <App />
    // },
    // {
    //   path:"/return", 
    //   element: <Return />
    // },
  ]);

  return (
    <StrictMode>
        <RouterProvider router={router} />
    </StrictMode>
  );
};

export default RouterComponent;
