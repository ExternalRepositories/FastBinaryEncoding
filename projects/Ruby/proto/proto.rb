# Automatically generated by the Fast Binary Encoding compiler, do not modify!
# https://github.com/chronoxor/FastBinaryEncoding

# rubocop:disable Lint/MissingCopEnableDirective
# rubocop:disable Lint/UnneededCopDisableDirective
# rubocop:disable Metrics/AbcSize
# rubocop:disable Metrics/ClassLength
# rubocop:disable Metrics/CyclomaticComplexity
# rubocop:disable Metrics/LineLength
# rubocop:disable Metrics/MethodLength
# rubocop:disable Metrics/PerceivedComplexity

require 'bigdecimal'
require 'uuidtools'

require_relative 'fbe'

module Proto

  module OrderSide
    class Enum
      include FBE::Enum

      define :buy, 0 + 0
      define :sell, 0 + 1

      def initialize(value = 0)
        @value = value
      end

      def to_i
        @value
      end

      def to_s
        if @value == Enum.buy
          return 'buy'
        end
        if @value == Enum.sell
          return 'sell'
        end
        '<unknown>'
      end
    end

    class << self
      attr_accessor :buy
      attr_accessor :sell
    end

    self.buy = Enum.new(Enum.buy)
    self.sell = Enum.new(Enum.sell)

    def self.new(value = 0)
      Enum.new(value)
    end
  end

  OrderSide.freeze

  # Fast Binary Encoding OrderSide field model class
  class FieldModelOrderSide < FBE::FieldModel
    def initialize(buffer, offset)
        super(buffer, offset)
    end

    # Get the field size
    def fbe_size
      1
    end

    # Get the value
    def get(defaults = OrderSide.new)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return defaults
      end

      OrderSide.new(read_byte(fbe_offset))
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return
      end

      write_byte(fbe_offset, value)
    end
  end

  # Fast Binary Encoding OrderSide final model class
  class FinalModelOrderSide < FBE::FinalModel
    def initialize(buffer, offset)
      super(buffer, offset)
    end

    # Get the allocation size
    # noinspection RubyUnusedLocalVariable
    def fbe_allocation_size(value)
      fbe_size
    end

    # Get the final size
    def fbe_size
      1
    end

    # Check if the value is valid
    def verify
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return Fixnum::MAX
      end

      fbe_size
    end

    # Get the value
    def get
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return [OrderSide.new, 0]
      end

      [OrderSide.new(read_byte(fbe_offset)), fbe_size]
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return 0
      end

      write_byte(fbe_offset, value)
      fbe_size
    end
  end

  module OrderType
    class Enum
      include FBE::Enum

      define :market, 0 + 0
      define :limit, 0 + 1
      define :stop, 0 + 2

      def initialize(value = 0)
        @value = value
      end

      def to_i
        @value
      end

      def to_s
        if @value == Enum.market
          return 'market'
        end
        if @value == Enum.limit
          return 'limit'
        end
        if @value == Enum.stop
          return 'stop'
        end
        '<unknown>'
      end
    end

    class << self
      attr_accessor :market
      attr_accessor :limit
      attr_accessor :stop
    end

    self.market = Enum.new(Enum.market)
    self.limit = Enum.new(Enum.limit)
    self.stop = Enum.new(Enum.stop)

    def self.new(value = 0)
      Enum.new(value)
    end
  end

  OrderType.freeze

  # Fast Binary Encoding OrderType field model class
  class FieldModelOrderType < FBE::FieldModel
    def initialize(buffer, offset)
        super(buffer, offset)
    end

    # Get the field size
    def fbe_size
      1
    end

    # Get the value
    def get(defaults = OrderType.new)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return defaults
      end

      OrderType.new(read_byte(fbe_offset))
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return
      end

      write_byte(fbe_offset, value)
    end
  end

  # Fast Binary Encoding OrderType final model class
  class FinalModelOrderType < FBE::FinalModel
    def initialize(buffer, offset)
      super(buffer, offset)
    end

    # Get the allocation size
    # noinspection RubyUnusedLocalVariable
    def fbe_allocation_size(value)
      fbe_size
    end

    # Get the final size
    def fbe_size
      1
    end

    # Check if the value is valid
    def verify
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return Fixnum::MAX
      end

      fbe_size
    end

    # Get the value
    def get
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return [OrderType.new, 0]
      end

      [OrderType.new(read_byte(fbe_offset)), fbe_size]
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return 0
      end

      write_byte(fbe_offset, value)
      fbe_size
    end
  end

  module State
    class Flags
      include FBE::Flags

      define :unknown, 0x00
      define :invalid, 0x01
      define :initialized, 0x02
      define :calculated, 0x04
      define :broken, 0x08
      define :good, Flags.value(:initialized)|Flags.value(:calculated)
      define :bad, Flags.value(:unknown)|Flags.value(:invalid)|Flags.value(:broken)

      def initialize(value = 0)
        @value = value
      end

      def to_i
        @value
      end

      def to_s
        result = ''
        first = true
        if (@value == Flags.unknown) && ((@value & Flags.unknown) == Flags.unknown)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'unknown'
        end
        if (@value == Flags.invalid) && ((@value & Flags.invalid) == Flags.invalid)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'invalid'
        end
        if (@value == Flags.initialized) && ((@value & Flags.initialized) == Flags.initialized)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'initialized'
        end
        if (@value == Flags.calculated) && ((@value & Flags.calculated) == Flags.calculated)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'calculated'
        end
        if (@value == Flags.broken) && ((@value & Flags.broken) == Flags.broken)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'broken'
        end
        if (@value == Flags.good) && ((@value & Flags.good) == Flags.good)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'good'
        end
        if (@value == Flags.bad) && ((@value & Flags.bad) == Flags.bad)
          if first
            # noinspection RubyUnusedLocalVariable
            first = false
          else
            result << '|'
          end
          result << 'bad'
        end
        result
      end
    end

    class << self
      attr_accessor :unknown
      attr_accessor :invalid
      attr_accessor :initialized
      attr_accessor :calculated
      attr_accessor :broken
      attr_accessor :good
      attr_accessor :bad
    end

    self.unknown = Flags.new(Flags.unknown)
    self.invalid = Flags.new(Flags.invalid)
    self.initialized = Flags.new(Flags.initialized)
    self.calculated = Flags.new(Flags.calculated)
    self.broken = Flags.new(Flags.broken)
    self.good = Flags.new(Flags.good)
    self.bad = Flags.new(Flags.bad)

    def self.new(value = 0)
      Flags.new(value)
    end
  end

  State.freeze

  # Fast Binary Encoding State field model class
  class FieldModelState < FBE::FieldModel
    def initialize(buffer, offset)
        super(buffer, offset)
    end

    # Get the field size
    def fbe_size
      1
    end

    # Get the value
    def get(defaults = State.new)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return defaults
      end

      State.new(read_byte(fbe_offset))
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return
      end

      write_byte(fbe_offset, value)
    end
  end

  # Fast Binary Encoding State final model class
  class FinalModelState < FBE::FinalModel
    def initialize(buffer, offset)
      super(buffer, offset)
    end

    # Get the allocation size
    # noinspection RubyUnusedLocalVariable
    def fbe_allocation_size(value)
      fbe_size
    end

    # Get the final size
    def fbe_size
      1
    end

    # Check if the value is valid
    def verify
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return Fixnum::MAX
      end

      fbe_size
    end

    # Get the value
    def get
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return [State.new, 0]
      end

      [State.new(read_byte(fbe_offset)), fbe_size]
    end

    # Set the value
    def set(value)
      if (@_buffer.offset + fbe_offset + fbe_size) > @_buffer.size
        return 0
      end

      write_byte(fbe_offset, value)
      fbe_size
    end
  end

end

# rubocop:enable all
