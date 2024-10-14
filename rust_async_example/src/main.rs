#[tokio::main]
async fn main() {
    let result = some_async_function().await;
    println!("Result: {:?}", result);
}

async fn some_async_function() -> i32 {
    42
}
