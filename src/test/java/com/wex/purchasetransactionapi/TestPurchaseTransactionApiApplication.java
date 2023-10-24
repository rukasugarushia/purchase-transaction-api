package com.wex.purchasetransactionapi;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.test.context.TestConfiguration;

@TestConfiguration(proxyBeanMethods = false)
public class TestPurchaseTransactionApiApplication {

	public static void main(String[] args) {
		SpringApplication.from(PurchaseTransactionApiApplication::main).with(TestPurchaseTransactionApiApplication.class).run(args);
	}

}
