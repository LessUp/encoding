use std::env;
use std::process;

pub trait FileProcessor {
    fn encode_file(&self, input_path: &str, output_path: &str) -> std::io::Result<()>;
    fn decode_file(&self, input_path: &str, output_path: &str) -> std::io::Result<()>;
}

pub fn run(_name: &str, processor: &dyn FileProcessor) {
    let args: Vec<String> = env::args().collect();

    if args.len() != 4 {
        eprintln!("Usage: {} encode|decode input output", args[0]);
        process::exit(1);
    }

    let mode = &args[1];
    let input_path = &args[2];
    let output_path = &args[3];

    let result = match mode.as_str() {
        "encode" => processor.encode_file(input_path, output_path),
        "decode" => processor.decode_file(input_path, output_path),
        _ => {
            eprintln!("unknown mode, expected encode or decode");
            process::exit(1);
        }
    };

    if let Err(e) = result {
        eprintln!("{}", e);
        process::exit(1);
    }
}
