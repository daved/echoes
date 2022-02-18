use std::io::Read;
use std::io::Write;
use std::net::{TcpListener, TcpStream};
use std::thread;

fn handle_client(i: u32, mut stream: TcpStream) {
    println!(
        "{}: {}:{} <-> {}:{}",
        i,
        stream.local_addr().unwrap().ip(),
        stream.local_addr().unwrap().port(),
        stream.peer_addr().unwrap().ip(),
        stream.peer_addr().unwrap().port(),
    );

    loop {
        let mut read = [0; 1028];
        match stream.read(&mut read) {
            Ok(n) => {
                if n == 0 {
                    // connection was closed
                    break;
                }
                stream.write(&read[0..n]).unwrap();
            }
            Err(err) => {
                std::panic::panic_any(err);
            }
        }
    }
}

fn main() {
    let listener = TcpListener::bind("127.0.0.1:25001").unwrap();
    let mut count = 0u32;

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                thread::spawn(move || {
                    handle_client(count, stream);
                });
            }
            Err(_) => {
                println!("Error");
            }
        }
        count += 1;
    }
}
