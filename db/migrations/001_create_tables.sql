-- Create database
CREATE DATABASE warehouse_management;
USE warehouse_management;


-- Table Master Supplier
CREATE TABLE MasterSupplier (
                                SupplierPK INT PRIMARY KEY AUTO_INCREMENT,
                                SupplierName NVARCHAR(255) NOT NULL
);

-- Table Master Customer
CREATE TABLE MasterCustomer (
                                CustomerPK INT PRIMARY KEY AUTO_INCREMENT,
                                CustomerName NVARCHAR(255) NOT NULL
);

-- Table Master Product
CREATE TABLE MasterProduct (
                               ProductPK INT PRIMARY KEY AUTO_INCREMENT,
                               ProductName NVARCHAR(255) NOT NULL
);

-- Table Master Warehouse
CREATE TABLE MasterWarehouse (
                                 WhsPK INT PRIMARY KEY AUTO_INCREMENT,
                                 WhsName NVARCHAR(255) NOT NULL
);

-- Table Transaksi PenerimaanBarangHeader
CREATE TABLE PenerimaanBarangHeader (
                                        TrxInPK INT PRIMARY KEY AUTO_INCREMENT,
                                        TrxInNo NVARCHAR(50) NOT NULL,
                                        WhsIdf INT NOT NULL,
                                        TrxInDate DATE NOT NULL,
                                        TrxInSuppIdf INT NOT NULL,
                                        TrxInNotes NVARCHAR(255),
                                        FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
                                        FOREIGN KEY (TrxInSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

-- Table Transaksi PenerimaanBarangDetail
CREATE TABLE PenerimaanBarangDetail (
                                        TrxInDPK INT PRIMARY KEY AUTO_INCREMENT,
                                        TrxInIDF INT NOT NULL,
                                        TrxInDProductIdf INT NOT NULL,
                                        TrxInDQtyDus INT NOT NULL,
                                        TrxInDQtyPcs INT NOT NULL,
                                        FOREIGN KEY (TrxInIDF) REFERENCES PenerimaanBarangHeader(TrxInPK),
                                        FOREIGN KEY (TrxInDProductIdf) REFERENCES MasterProduct(ProductPK)
);

-- Table Transaksi PengeluaranBarangHeader
CREATE TABLE PengeluaranBarangHeader (
                                         TrxOutPK INT PRIMARY KEY AUTO_INCREMENT,
                                         TrxOutNo NVARCHAR(50) NOT NULL,
                                         WhsIdf INT NOT NULL,
                                         TrxOutDate DATE NOT NULL,
                                         TrxOutSuppIdf INT NOT NULL,
                                         TrxOutNotes NVARCHAR(255),
                                         FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
                                         FOREIGN KEY (TrxOutSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

-- Table Transaksi PengeluaranBarangDetail
CREATE TABLE PengeluaranBarangDetail (
                                         TrxOutDPK INT PRIMARY KEY AUTO_INCREMENT,
                                         TrxOutIDF INT NOT NULL,
                                         TrxOutDProductIdf INT NOT NULL,
                                         TrxOutDQtyDus INT NOT NULL,
                                         TrxOutDQtyPcs INT NOT NULL,
                                         FOREIGN KEY (TrxOutIDF) REFERENCES PengeluaranBarangHeader(TrxOutPK),
                                         FOREIGN KEY (TrxOutDProductIdf) REFERENCES MasterProduct(ProductPK)
);