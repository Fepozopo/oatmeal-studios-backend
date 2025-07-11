// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.assignPlanogramToLocationStmt, err = db.PrepareContext(ctx, assignPlanogramToLocation); err != nil {
		return nil, fmt.Errorf("error preparing query AssignPlanogramToLocation: %w", err)
	}
	if q.createCustomerStmt, err = db.PrepareContext(ctx, createCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCustomer: %w", err)
	}
	if q.createCustomerLocationStmt, err = db.PrepareContext(ctx, createCustomerLocation); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCustomerLocation: %w", err)
	}
	if q.createInvoiceStmt, err = db.PrepareContext(ctx, createInvoice); err != nil {
		return nil, fmt.Errorf("error preparing query CreateInvoice: %w", err)
	}
	if q.createOrderStmt, err = db.PrepareContext(ctx, createOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrder: %w", err)
	}
	if q.createOrderItemStmt, err = db.PrepareContext(ctx, createOrderItem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrderItem: %w", err)
	}
	if q.createPlanogramStmt, err = db.PrepareContext(ctx, createPlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlanogram: %w", err)
	}
	if q.createPlanogramPocketStmt, err = db.PrepareContext(ctx, createPlanogramPocket); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlanogramPocket: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.createRefreshTokenStmt, err = db.PrepareContext(ctx, createRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRefreshToken: %w", err)
	}
	if q.createSalesRepStmt, err = db.PrepareContext(ctx, createSalesRep); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSalesRep: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteAllUsersStmt, err = db.PrepareContext(ctx, deleteAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAllUsers: %w", err)
	}
	if q.deleteCustomerStmt, err = db.PrepareContext(ctx, deleteCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCustomer: %w", err)
	}
	if q.deleteCustomerLocationStmt, err = db.PrepareContext(ctx, deleteCustomerLocation); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCustomerLocation: %w", err)
	}
	if q.deleteInvoiceStmt, err = db.PrepareContext(ctx, deleteInvoice); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteInvoice: %w", err)
	}
	if q.deleteOrderStmt, err = db.PrepareContext(ctx, deleteOrder); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOrder: %w", err)
	}
	if q.deleteOrderItemStmt, err = db.PrepareContext(ctx, deleteOrderItem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOrderItem: %w", err)
	}
	if q.deletePlanogramStmt, err = db.PrepareContext(ctx, deletePlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePlanogram: %w", err)
	}
	if q.deletePlanogramPocketStmt, err = db.PrepareContext(ctx, deletePlanogramPocket); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePlanogramPocket: %w", err)
	}
	if q.deleteProductStmt, err = db.PrepareContext(ctx, deleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProduct: %w", err)
	}
	if q.deleteSalesRepStmt, err = db.PrepareContext(ctx, deleteSalesRep); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSalesRep: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAllCurrentInventoryStmt, err = db.PrepareContext(ctx, getAllCurrentInventory); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCurrentInventory: %w", err)
	}
	if q.getCurrentInventoryStmt, err = db.PrepareContext(ctx, getCurrentInventory); err != nil {
		return nil, fmt.Errorf("error preparing query GetCurrentInventory: %w", err)
	}
	if q.getCustomerStmt, err = db.PrepareContext(ctx, getCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query GetCustomer: %w", err)
	}
	if q.getCustomerLocationByIDStmt, err = db.PrepareContext(ctx, getCustomerLocationByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCustomerLocationByID: %w", err)
	}
	if q.getInventoryChangesByDayStmt, err = db.PrepareContext(ctx, getInventoryChangesByDay); err != nil {
		return nil, fmt.Errorf("error preparing query GetInventoryChangesByDay: %w", err)
	}
	if q.getInvoiceStmt, err = db.PrepareContext(ctx, getInvoice); err != nil {
		return nil, fmt.Errorf("error preparing query GetInvoice: %w", err)
	}
	if q.getInvoicesByOrderStmt, err = db.PrepareContext(ctx, getInvoicesByOrder); err != nil {
		return nil, fmt.Errorf("error preparing query GetInvoicesByOrder: %w", err)
	}
	if q.getOrderStmt, err = db.PrepareContext(ctx, getOrder); err != nil {
		return nil, fmt.Errorf("error preparing query GetOrder: %w", err)
	}
	if q.getOrderItemStmt, err = db.PrepareContext(ctx, getOrderItem); err != nil {
		return nil, fmt.Errorf("error preparing query GetOrderItem: %w", err)
	}
	if q.getPlanogramStmt, err = db.PrepareContext(ctx, getPlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlanogram: %w", err)
	}
	if q.getPlanogramPocketStmt, err = db.PrepareContext(ctx, getPlanogramPocket); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlanogramPocket: %w", err)
	}
	if q.getPlanogramPocketByNumberStmt, err = db.PrepareContext(ctx, getPlanogramPocketByNumber); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlanogramPocketByNumber: %w", err)
	}
	if q.getPlanogramsByLocationStmt, err = db.PrepareContext(ctx, getPlanogramsByLocation); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlanogramsByLocation: %w", err)
	}
	if q.getProductByIDStmt, err = db.PrepareContext(ctx, getProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductByID: %w", err)
	}
	if q.getProductBySKUStmt, err = db.PrepareContext(ctx, getProductBySKU); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductBySKU: %w", err)
	}
	if q.getRefreshTokenStmt, err = db.PrepareContext(ctx, getRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetRefreshToken: %w", err)
	}
	if q.getSalesRepStmt, err = db.PrepareContext(ctx, getSalesRep); err != nil {
		return nil, fmt.Errorf("error preparing query GetSalesRep: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIDStmt, err = db.PrepareContext(ctx, getUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByID: %w", err)
	}
	if q.getUserFromRefreshTokenStmt, err = db.PrepareContext(ctx, getUserFromRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserFromRefreshToken: %w", err)
	}
	if q.insertInventoryTransactionStmt, err = db.PrepareContext(ctx, insertInventoryTransaction); err != nil {
		return nil, fmt.Errorf("error preparing query InsertInventoryTransaction: %w", err)
	}
	if q.listCustomerLocationsByCustomerStmt, err = db.PrepareContext(ctx, listCustomerLocationsByCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query ListCustomerLocationsByCustomer: %w", err)
	}
	if q.listCustomersStmt, err = db.PrepareContext(ctx, listCustomers); err != nil {
		return nil, fmt.Errorf("error preparing query ListCustomers: %w", err)
	}
	if q.listInventoryTransactionsStmt, err = db.PrepareContext(ctx, listInventoryTransactions); err != nil {
		return nil, fmt.Errorf("error preparing query ListInventoryTransactions: %w", err)
	}
	if q.listInvoicesByCustomerStmt, err = db.PrepareContext(ctx, listInvoicesByCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query ListInvoicesByCustomer: %w", err)
	}
	if q.listInvoicesByCustomerLocationStmt, err = db.PrepareContext(ctx, listInvoicesByCustomerLocation); err != nil {
		return nil, fmt.Errorf("error preparing query ListInvoicesByCustomerLocation: %w", err)
	}
	if q.listLocationsByPlanogramStmt, err = db.PrepareContext(ctx, listLocationsByPlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query ListLocationsByPlanogram: %w", err)
	}
	if q.listOrderItemsByOrderIDStmt, err = db.PrepareContext(ctx, listOrderItemsByOrderID); err != nil {
		return nil, fmt.Errorf("error preparing query ListOrderItemsByOrderID: %w", err)
	}
	if q.listOrderItemsBySKUStmt, err = db.PrepareContext(ctx, listOrderItemsBySKU); err != nil {
		return nil, fmt.Errorf("error preparing query ListOrderItemsBySKU: %w", err)
	}
	if q.listOrdersByCustomerStmt, err = db.PrepareContext(ctx, listOrdersByCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query ListOrdersByCustomer: %w", err)
	}
	if q.listOrdersOpenStmt, err = db.PrepareContext(ctx, listOrdersOpen); err != nil {
		return nil, fmt.Errorf("error preparing query ListOrdersOpen: %w", err)
	}
	if q.listPlanogramsStmt, err = db.PrepareContext(ctx, listPlanograms); err != nil {
		return nil, fmt.Errorf("error preparing query ListPlanograms: %w", err)
	}
	if q.listPocketsForPlanogramStmt, err = db.PrepareContext(ctx, listPocketsForPlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query ListPocketsForPlanogram: %w", err)
	}
	if q.listProductsByArtistStmt, err = db.PrepareContext(ctx, listProductsByArtist); err != nil {
		return nil, fmt.Errorf("error preparing query ListProductsByArtist: %w", err)
	}
	if q.listProductsByCategoryStmt, err = db.PrepareContext(ctx, listProductsByCategory); err != nil {
		return nil, fmt.Errorf("error preparing query ListProductsByCategory: %w", err)
	}
	if q.listProductsByStatusStmt, err = db.PrepareContext(ctx, listProductsByStatus); err != nil {
		return nil, fmt.Errorf("error preparing query ListProductsByStatus: %w", err)
	}
	if q.listProductsByTypeStmt, err = db.PrepareContext(ctx, listProductsByType); err != nil {
		return nil, fmt.Errorf("error preparing query ListProductsByType: %w", err)
	}
	if q.listSalesRepsStmt, err = db.PrepareContext(ctx, listSalesReps); err != nil {
		return nil, fmt.Errorf("error preparing query ListSalesReps: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.reassignPlanogramToLocationStmt, err = db.PrepareContext(ctx, reassignPlanogramToLocation); err != nil {
		return nil, fmt.Errorf("error preparing query ReassignPlanogramToLocation: %w", err)
	}
	if q.removePlanogramFromLocationStmt, err = db.PrepareContext(ctx, removePlanogramFromLocation); err != nil {
		return nil, fmt.Errorf("error preparing query RemovePlanogramFromLocation: %w", err)
	}
	if q.revokeRefreshTokenStmt, err = db.PrepareContext(ctx, revokeRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query RevokeRefreshToken: %w", err)
	}
	if q.updateCustomerStmt, err = db.PrepareContext(ctx, updateCustomer); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCustomer: %w", err)
	}
	if q.updateCustomerLocationStmt, err = db.PrepareContext(ctx, updateCustomerLocation); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCustomerLocation: %w", err)
	}
	if q.updateInvoiceStmt, err = db.PrepareContext(ctx, updateInvoice); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateInvoice: %w", err)
	}
	if q.updateOrderStmt, err = db.PrepareContext(ctx, updateOrder); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOrder: %w", err)
	}
	if q.updateOrderItemStmt, err = db.PrepareContext(ctx, updateOrderItem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOrderItem: %w", err)
	}
	if q.updatePlanogramStmt, err = db.PrepareContext(ctx, updatePlanogram); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePlanogram: %w", err)
	}
	if q.updatePlanogramPocketStmt, err = db.PrepareContext(ctx, updatePlanogramPocket); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePlanogramPocket: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	if q.updateSalesRepStmt, err = db.PrepareContext(ctx, updateSalesRep); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSalesRep: %w", err)
	}
	if q.updateUserNameStmt, err = db.PrepareContext(ctx, updateUserName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserName: %w", err)
	}
	if q.updateUserPasswordStmt, err = db.PrepareContext(ctx, updateUserPassword); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserPassword: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.assignPlanogramToLocationStmt != nil {
		if cerr := q.assignPlanogramToLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing assignPlanogramToLocationStmt: %w", cerr)
		}
	}
	if q.createCustomerStmt != nil {
		if cerr := q.createCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCustomerStmt: %w", cerr)
		}
	}
	if q.createCustomerLocationStmt != nil {
		if cerr := q.createCustomerLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCustomerLocationStmt: %w", cerr)
		}
	}
	if q.createInvoiceStmt != nil {
		if cerr := q.createInvoiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createInvoiceStmt: %w", cerr)
		}
	}
	if q.createOrderStmt != nil {
		if cerr := q.createOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderStmt: %w", cerr)
		}
	}
	if q.createOrderItemStmt != nil {
		if cerr := q.createOrderItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderItemStmt: %w", cerr)
		}
	}
	if q.createPlanogramStmt != nil {
		if cerr := q.createPlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlanogramStmt: %w", cerr)
		}
	}
	if q.createPlanogramPocketStmt != nil {
		if cerr := q.createPlanogramPocketStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlanogramPocketStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.createRefreshTokenStmt != nil {
		if cerr := q.createRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRefreshTokenStmt: %w", cerr)
		}
	}
	if q.createSalesRepStmt != nil {
		if cerr := q.createSalesRepStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSalesRepStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteAllUsersStmt != nil {
		if cerr := q.deleteAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAllUsersStmt: %w", cerr)
		}
	}
	if q.deleteCustomerStmt != nil {
		if cerr := q.deleteCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCustomerStmt: %w", cerr)
		}
	}
	if q.deleteCustomerLocationStmt != nil {
		if cerr := q.deleteCustomerLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCustomerLocationStmt: %w", cerr)
		}
	}
	if q.deleteInvoiceStmt != nil {
		if cerr := q.deleteInvoiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteInvoiceStmt: %w", cerr)
		}
	}
	if q.deleteOrderStmt != nil {
		if cerr := q.deleteOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOrderStmt: %w", cerr)
		}
	}
	if q.deleteOrderItemStmt != nil {
		if cerr := q.deleteOrderItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOrderItemStmt: %w", cerr)
		}
	}
	if q.deletePlanogramStmt != nil {
		if cerr := q.deletePlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePlanogramStmt: %w", cerr)
		}
	}
	if q.deletePlanogramPocketStmt != nil {
		if cerr := q.deletePlanogramPocketStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePlanogramPocketStmt: %w", cerr)
		}
	}
	if q.deleteProductStmt != nil {
		if cerr := q.deleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductStmt: %w", cerr)
		}
	}
	if q.deleteSalesRepStmt != nil {
		if cerr := q.deleteSalesRepStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSalesRepStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAllCurrentInventoryStmt != nil {
		if cerr := q.getAllCurrentInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCurrentInventoryStmt: %w", cerr)
		}
	}
	if q.getCurrentInventoryStmt != nil {
		if cerr := q.getCurrentInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCurrentInventoryStmt: %w", cerr)
		}
	}
	if q.getCustomerStmt != nil {
		if cerr := q.getCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCustomerStmt: %w", cerr)
		}
	}
	if q.getCustomerLocationByIDStmt != nil {
		if cerr := q.getCustomerLocationByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCustomerLocationByIDStmt: %w", cerr)
		}
	}
	if q.getInventoryChangesByDayStmt != nil {
		if cerr := q.getInventoryChangesByDayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInventoryChangesByDayStmt: %w", cerr)
		}
	}
	if q.getInvoiceStmt != nil {
		if cerr := q.getInvoiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInvoiceStmt: %w", cerr)
		}
	}
	if q.getInvoicesByOrderStmt != nil {
		if cerr := q.getInvoicesByOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInvoicesByOrderStmt: %w", cerr)
		}
	}
	if q.getOrderStmt != nil {
		if cerr := q.getOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOrderStmt: %w", cerr)
		}
	}
	if q.getOrderItemStmt != nil {
		if cerr := q.getOrderItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOrderItemStmt: %w", cerr)
		}
	}
	if q.getPlanogramStmt != nil {
		if cerr := q.getPlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlanogramStmt: %w", cerr)
		}
	}
	if q.getPlanogramPocketStmt != nil {
		if cerr := q.getPlanogramPocketStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlanogramPocketStmt: %w", cerr)
		}
	}
	if q.getPlanogramPocketByNumberStmt != nil {
		if cerr := q.getPlanogramPocketByNumberStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlanogramPocketByNumberStmt: %w", cerr)
		}
	}
	if q.getPlanogramsByLocationStmt != nil {
		if cerr := q.getPlanogramsByLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlanogramsByLocationStmt: %w", cerr)
		}
	}
	if q.getProductByIDStmt != nil {
		if cerr := q.getProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductByIDStmt: %w", cerr)
		}
	}
	if q.getProductBySKUStmt != nil {
		if cerr := q.getProductBySKUStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductBySKUStmt: %w", cerr)
		}
	}
	if q.getRefreshTokenStmt != nil {
		if cerr := q.getRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRefreshTokenStmt: %w", cerr)
		}
	}
	if q.getSalesRepStmt != nil {
		if cerr := q.getSalesRepStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSalesRepStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIDStmt != nil {
		if cerr := q.getUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIDStmt: %w", cerr)
		}
	}
	if q.getUserFromRefreshTokenStmt != nil {
		if cerr := q.getUserFromRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserFromRefreshTokenStmt: %w", cerr)
		}
	}
	if q.insertInventoryTransactionStmt != nil {
		if cerr := q.insertInventoryTransactionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertInventoryTransactionStmt: %w", cerr)
		}
	}
	if q.listCustomerLocationsByCustomerStmt != nil {
		if cerr := q.listCustomerLocationsByCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCustomerLocationsByCustomerStmt: %w", cerr)
		}
	}
	if q.listCustomersStmt != nil {
		if cerr := q.listCustomersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCustomersStmt: %w", cerr)
		}
	}
	if q.listInventoryTransactionsStmt != nil {
		if cerr := q.listInventoryTransactionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInventoryTransactionsStmt: %w", cerr)
		}
	}
	if q.listInvoicesByCustomerStmt != nil {
		if cerr := q.listInvoicesByCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInvoicesByCustomerStmt: %w", cerr)
		}
	}
	if q.listInvoicesByCustomerLocationStmt != nil {
		if cerr := q.listInvoicesByCustomerLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInvoicesByCustomerLocationStmt: %w", cerr)
		}
	}
	if q.listLocationsByPlanogramStmt != nil {
		if cerr := q.listLocationsByPlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listLocationsByPlanogramStmt: %w", cerr)
		}
	}
	if q.listOrderItemsByOrderIDStmt != nil {
		if cerr := q.listOrderItemsByOrderIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listOrderItemsByOrderIDStmt: %w", cerr)
		}
	}
	if q.listOrderItemsBySKUStmt != nil {
		if cerr := q.listOrderItemsBySKUStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listOrderItemsBySKUStmt: %w", cerr)
		}
	}
	if q.listOrdersByCustomerStmt != nil {
		if cerr := q.listOrdersByCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listOrdersByCustomerStmt: %w", cerr)
		}
	}
	if q.listOrdersOpenStmt != nil {
		if cerr := q.listOrdersOpenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listOrdersOpenStmt: %w", cerr)
		}
	}
	if q.listPlanogramsStmt != nil {
		if cerr := q.listPlanogramsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPlanogramsStmt: %w", cerr)
		}
	}
	if q.listPocketsForPlanogramStmt != nil {
		if cerr := q.listPocketsForPlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPocketsForPlanogramStmt: %w", cerr)
		}
	}
	if q.listProductsByArtistStmt != nil {
		if cerr := q.listProductsByArtistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductsByArtistStmt: %w", cerr)
		}
	}
	if q.listProductsByCategoryStmt != nil {
		if cerr := q.listProductsByCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductsByCategoryStmt: %w", cerr)
		}
	}
	if q.listProductsByStatusStmt != nil {
		if cerr := q.listProductsByStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductsByStatusStmt: %w", cerr)
		}
	}
	if q.listProductsByTypeStmt != nil {
		if cerr := q.listProductsByTypeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductsByTypeStmt: %w", cerr)
		}
	}
	if q.listSalesRepsStmt != nil {
		if cerr := q.listSalesRepsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSalesRepsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.reassignPlanogramToLocationStmt != nil {
		if cerr := q.reassignPlanogramToLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing reassignPlanogramToLocationStmt: %w", cerr)
		}
	}
	if q.removePlanogramFromLocationStmt != nil {
		if cerr := q.removePlanogramFromLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removePlanogramFromLocationStmt: %w", cerr)
		}
	}
	if q.revokeRefreshTokenStmt != nil {
		if cerr := q.revokeRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing revokeRefreshTokenStmt: %w", cerr)
		}
	}
	if q.updateCustomerStmt != nil {
		if cerr := q.updateCustomerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCustomerStmt: %w", cerr)
		}
	}
	if q.updateCustomerLocationStmt != nil {
		if cerr := q.updateCustomerLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCustomerLocationStmt: %w", cerr)
		}
	}
	if q.updateInvoiceStmt != nil {
		if cerr := q.updateInvoiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateInvoiceStmt: %w", cerr)
		}
	}
	if q.updateOrderStmt != nil {
		if cerr := q.updateOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOrderStmt: %w", cerr)
		}
	}
	if q.updateOrderItemStmt != nil {
		if cerr := q.updateOrderItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOrderItemStmt: %w", cerr)
		}
	}
	if q.updatePlanogramStmt != nil {
		if cerr := q.updatePlanogramStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePlanogramStmt: %w", cerr)
		}
	}
	if q.updatePlanogramPocketStmt != nil {
		if cerr := q.updatePlanogramPocketStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePlanogramPocketStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
		}
	}
	if q.updateSalesRepStmt != nil {
		if cerr := q.updateSalesRepStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSalesRepStmt: %w", cerr)
		}
	}
	if q.updateUserNameStmt != nil {
		if cerr := q.updateUserNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserNameStmt: %w", cerr)
		}
	}
	if q.updateUserPasswordStmt != nil {
		if cerr := q.updateUserPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserPasswordStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                  DBTX
	tx                                  *sql.Tx
	assignPlanogramToLocationStmt       *sql.Stmt
	createCustomerStmt                  *sql.Stmt
	createCustomerLocationStmt          *sql.Stmt
	createInvoiceStmt                   *sql.Stmt
	createOrderStmt                     *sql.Stmt
	createOrderItemStmt                 *sql.Stmt
	createPlanogramStmt                 *sql.Stmt
	createPlanogramPocketStmt           *sql.Stmt
	createProductStmt                   *sql.Stmt
	createRefreshTokenStmt              *sql.Stmt
	createSalesRepStmt                  *sql.Stmt
	createUserStmt                      *sql.Stmt
	deleteAllUsersStmt                  *sql.Stmt
	deleteCustomerStmt                  *sql.Stmt
	deleteCustomerLocationStmt          *sql.Stmt
	deleteInvoiceStmt                   *sql.Stmt
	deleteOrderStmt                     *sql.Stmt
	deleteOrderItemStmt                 *sql.Stmt
	deletePlanogramStmt                 *sql.Stmt
	deletePlanogramPocketStmt           *sql.Stmt
	deleteProductStmt                   *sql.Stmt
	deleteSalesRepStmt                  *sql.Stmt
	deleteUserStmt                      *sql.Stmt
	getAllCurrentInventoryStmt          *sql.Stmt
	getCurrentInventoryStmt             *sql.Stmt
	getCustomerStmt                     *sql.Stmt
	getCustomerLocationByIDStmt         *sql.Stmt
	getInventoryChangesByDayStmt        *sql.Stmt
	getInvoiceStmt                      *sql.Stmt
	getInvoicesByOrderStmt              *sql.Stmt
	getOrderStmt                        *sql.Stmt
	getOrderItemStmt                    *sql.Stmt
	getPlanogramStmt                    *sql.Stmt
	getPlanogramPocketStmt              *sql.Stmt
	getPlanogramPocketByNumberStmt      *sql.Stmt
	getPlanogramsByLocationStmt         *sql.Stmt
	getProductByIDStmt                  *sql.Stmt
	getProductBySKUStmt                 *sql.Stmt
	getRefreshTokenStmt                 *sql.Stmt
	getSalesRepStmt                     *sql.Stmt
	getUserByEmailStmt                  *sql.Stmt
	getUserByIDStmt                     *sql.Stmt
	getUserFromRefreshTokenStmt         *sql.Stmt
	insertInventoryTransactionStmt      *sql.Stmt
	listCustomerLocationsByCustomerStmt *sql.Stmt
	listCustomersStmt                   *sql.Stmt
	listInventoryTransactionsStmt       *sql.Stmt
	listInvoicesByCustomerStmt          *sql.Stmt
	listInvoicesByCustomerLocationStmt  *sql.Stmt
	listLocationsByPlanogramStmt        *sql.Stmt
	listOrderItemsByOrderIDStmt         *sql.Stmt
	listOrderItemsBySKUStmt             *sql.Stmt
	listOrdersByCustomerStmt            *sql.Stmt
	listOrdersOpenStmt                  *sql.Stmt
	listPlanogramsStmt                  *sql.Stmt
	listPocketsForPlanogramStmt         *sql.Stmt
	listProductsByArtistStmt            *sql.Stmt
	listProductsByCategoryStmt          *sql.Stmt
	listProductsByStatusStmt            *sql.Stmt
	listProductsByTypeStmt              *sql.Stmt
	listSalesRepsStmt                   *sql.Stmt
	listUsersStmt                       *sql.Stmt
	reassignPlanogramToLocationStmt     *sql.Stmt
	removePlanogramFromLocationStmt     *sql.Stmt
	revokeRefreshTokenStmt              *sql.Stmt
	updateCustomerStmt                  *sql.Stmt
	updateCustomerLocationStmt          *sql.Stmt
	updateInvoiceStmt                   *sql.Stmt
	updateOrderStmt                     *sql.Stmt
	updateOrderItemStmt                 *sql.Stmt
	updatePlanogramStmt                 *sql.Stmt
	updatePlanogramPocketStmt           *sql.Stmt
	updateProductStmt                   *sql.Stmt
	updateSalesRepStmt                  *sql.Stmt
	updateUserNameStmt                  *sql.Stmt
	updateUserPasswordStmt              *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                  tx,
		tx:                                  tx,
		assignPlanogramToLocationStmt:       q.assignPlanogramToLocationStmt,
		createCustomerStmt:                  q.createCustomerStmt,
		createCustomerLocationStmt:          q.createCustomerLocationStmt,
		createInvoiceStmt:                   q.createInvoiceStmt,
		createOrderStmt:                     q.createOrderStmt,
		createOrderItemStmt:                 q.createOrderItemStmt,
		createPlanogramStmt:                 q.createPlanogramStmt,
		createPlanogramPocketStmt:           q.createPlanogramPocketStmt,
		createProductStmt:                   q.createProductStmt,
		createRefreshTokenStmt:              q.createRefreshTokenStmt,
		createSalesRepStmt:                  q.createSalesRepStmt,
		createUserStmt:                      q.createUserStmt,
		deleteAllUsersStmt:                  q.deleteAllUsersStmt,
		deleteCustomerStmt:                  q.deleteCustomerStmt,
		deleteCustomerLocationStmt:          q.deleteCustomerLocationStmt,
		deleteInvoiceStmt:                   q.deleteInvoiceStmt,
		deleteOrderStmt:                     q.deleteOrderStmt,
		deleteOrderItemStmt:                 q.deleteOrderItemStmt,
		deletePlanogramStmt:                 q.deletePlanogramStmt,
		deletePlanogramPocketStmt:           q.deletePlanogramPocketStmt,
		deleteProductStmt:                   q.deleteProductStmt,
		deleteSalesRepStmt:                  q.deleteSalesRepStmt,
		deleteUserStmt:                      q.deleteUserStmt,
		getAllCurrentInventoryStmt:          q.getAllCurrentInventoryStmt,
		getCurrentInventoryStmt:             q.getCurrentInventoryStmt,
		getCustomerStmt:                     q.getCustomerStmt,
		getCustomerLocationByIDStmt:         q.getCustomerLocationByIDStmt,
		getInventoryChangesByDayStmt:        q.getInventoryChangesByDayStmt,
		getInvoiceStmt:                      q.getInvoiceStmt,
		getInvoicesByOrderStmt:              q.getInvoicesByOrderStmt,
		getOrderStmt:                        q.getOrderStmt,
		getOrderItemStmt:                    q.getOrderItemStmt,
		getPlanogramStmt:                    q.getPlanogramStmt,
		getPlanogramPocketStmt:              q.getPlanogramPocketStmt,
		getPlanogramPocketByNumberStmt:      q.getPlanogramPocketByNumberStmt,
		getPlanogramsByLocationStmt:         q.getPlanogramsByLocationStmt,
		getProductByIDStmt:                  q.getProductByIDStmt,
		getProductBySKUStmt:                 q.getProductBySKUStmt,
		getRefreshTokenStmt:                 q.getRefreshTokenStmt,
		getSalesRepStmt:                     q.getSalesRepStmt,
		getUserByEmailStmt:                  q.getUserByEmailStmt,
		getUserByIDStmt:                     q.getUserByIDStmt,
		getUserFromRefreshTokenStmt:         q.getUserFromRefreshTokenStmt,
		insertInventoryTransactionStmt:      q.insertInventoryTransactionStmt,
		listCustomerLocationsByCustomerStmt: q.listCustomerLocationsByCustomerStmt,
		listCustomersStmt:                   q.listCustomersStmt,
		listInventoryTransactionsStmt:       q.listInventoryTransactionsStmt,
		listInvoicesByCustomerStmt:          q.listInvoicesByCustomerStmt,
		listInvoicesByCustomerLocationStmt:  q.listInvoicesByCustomerLocationStmt,
		listLocationsByPlanogramStmt:        q.listLocationsByPlanogramStmt,
		listOrderItemsByOrderIDStmt:         q.listOrderItemsByOrderIDStmt,
		listOrderItemsBySKUStmt:             q.listOrderItemsBySKUStmt,
		listOrdersByCustomerStmt:            q.listOrdersByCustomerStmt,
		listOrdersOpenStmt:                  q.listOrdersOpenStmt,
		listPlanogramsStmt:                  q.listPlanogramsStmt,
		listPocketsForPlanogramStmt:         q.listPocketsForPlanogramStmt,
		listProductsByArtistStmt:            q.listProductsByArtistStmt,
		listProductsByCategoryStmt:          q.listProductsByCategoryStmt,
		listProductsByStatusStmt:            q.listProductsByStatusStmt,
		listProductsByTypeStmt:              q.listProductsByTypeStmt,
		listSalesRepsStmt:                   q.listSalesRepsStmt,
		listUsersStmt:                       q.listUsersStmt,
		reassignPlanogramToLocationStmt:     q.reassignPlanogramToLocationStmt,
		removePlanogramFromLocationStmt:     q.removePlanogramFromLocationStmt,
		revokeRefreshTokenStmt:              q.revokeRefreshTokenStmt,
		updateCustomerStmt:                  q.updateCustomerStmt,
		updateCustomerLocationStmt:          q.updateCustomerLocationStmt,
		updateInvoiceStmt:                   q.updateInvoiceStmt,
		updateOrderStmt:                     q.updateOrderStmt,
		updateOrderItemStmt:                 q.updateOrderItemStmt,
		updatePlanogramStmt:                 q.updatePlanogramStmt,
		updatePlanogramPocketStmt:           q.updatePlanogramPocketStmt,
		updateProductStmt:                   q.updateProductStmt,
		updateSalesRepStmt:                  q.updateSalesRepStmt,
		updateUserNameStmt:                  q.updateUserNameStmt,
		updateUserPasswordStmt:              q.updateUserPasswordStmt,
	}
}
